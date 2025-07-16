<script lang="ts">
    import Header from "$lib/components/Header.svelte";
    import { goto } from "$app/navigation";
    import { registerUser } from "$lib/stores/auth";    
    
    let email = $state("");
    let password = $state("");
    let confirmPassword = $state("");
    let username = $state("");
    let firstName = $state("");
    let lastName = $state("");
    let fullBirthdate = $state("");
    let gender = $state("");
    let address = $state("");
    let phone = $state("");

    // New variables for birthdate
    let birthdateDay = $state("");
    let birthdateMonth = $state("");
    let birthdateYear = $state("");

    let error = $state("");

    async function handleRegister() {
        error = "";
        const fullBirthdate = `${birthdateYear}-${birthdateMonth}-${birthdateDay}`;
        if (password !== confirmPassword) {
            error = "Passwords do not match";
            return;
        }
        if (!username || !firstName || !lastName || !fullBirthdate || !gender || !address || !phone) {
            error = "All fields are required.";
            return;
        }
        try {
            await registerUser( email, password, username, firstName, lastName, fullBirthdate, gender, address, phone );
            goto("/");
        } catch (err) {
            console.error("Registration failed:", err);
            error = err.message || "Registration failed. Please try again.";
        }
    }
</script>

<Header />
<form class="register" onsubmit={handleRegister}>
    <h1>Registration</h1>
    <fieldset class="row1">
        <legend>Account Details
        </legend>
        <p>
            <label for="email">Email *</label>
            <input type="email" bind:value={email} required/>
            <label for="repeat-email">Repeat email *</label>
            <input type="text"/>
        </p>
        <p>
            <label for="password">Password*</label>
            <input type="password" bind:value={password} required/>
            <label for="repeat-password">Repeat Password*</label>
            <input type="password" bind:value={confirmPassword} required />
            <label for="obinfo" class="obinfo">* obligatory fields</label>
        </p>
        <p>
            <label for="username">Username*</label>
            <input type="text" bind:value={username} required />
        </p>
    </fieldset>
    <fieldset class="row2">
        <legend>Personal Details</legend>
        <p>
            <label for="firstName">First Name *</label>
            <input type="text" class="long" bind:value={firstName} required />
        </p>
        <p>
            <label for="lastName">Last Name *</label>
            <input type="text" class="long" bind:value={lastName} required />
        </p>
        <p>
            <label for="phone">Phone *</label>
            <input type="text" maxlength="10" bind:value={phone} required/>
        </p>
        <p>
            <label for="address" class="optional">Address</label>
            <input type="text" class="long" bind:value={address} required/>
        </p>
    </fieldset>
    <fieldset class="row3">
        <legend>Further Information
        </legend>
        <p>
            <label for="gender">Gender *</label>
            <input type="radio" name="gender" value="male" bind:group={gender} required/>
            <label for="male-gender" class="gender">Male</label>
            <input type="radio" name="gender" value="female" bind:group={gender} required />
            <label for="female-gender" class="gender">Female</label>
        </p>
        <p>
            <label for="birthdate">Birthdate *</label>
            <select class="date" bind:value={birthdateDay} required>
                <option value="">Day</option><option value="1">01</option>
                <option value="2">02</option>
                <option value="3">03</option>
                <option value="4">04</option>
                <option value="5">05</option>
                <option value="6">06</option>
                <option value="7">07</option>
                <option value="8">08</option>
                <option value="9">09</option>
                <option value="10">10</option>
                <option value="11">11</option>
                <option value="12">12</option>
                <option value="13">13</option>
                <option value="14">14</option>
                <option value="15">15</option>
                <option value="16">16</option>
                <option value="17">17</option>
                <option value="18">18</option>
                <option value="19">19</option>
                <option value="20">20</option>
                <option value="21">21</option>
                <option value="22">22</option>
                <option value="23">23</option>
                <option value="24">24</option>
                <option value="25">25</option>
                <option value="26">26</option>
                <option value="27">27</option>
                <option value="28">28</option>
                <option value="29">29</option>
                <option value="30">30</option>
                <option value="31">31</option>
            </select>
            <select bind:value={birthdateMonth} required>
                <option value="">Month</option> <option value="1">January</option> 
                <option value="2">February</option>
                <option value="3">March</option>
                <option value="4">April</option>
                <option value="5">May</option>
                <option value="6">June</option>
                <option value="7">July</option>
                <option value="8">August</option>
                <option value="9">September</option>
                <option value="10">October</option>
                <option value="11">November</option>
                <option value="12">December</option>
            </select>
            <input class="year" type="text" size="4" maxlength="4" bind:value={birthdateYear} placeholder="Year"/>e.g 1976
        </p>
    </fieldset>
    <div class="register-btn">
        <button class="button" type="submit">Register &raquo;</button>
        <p class="message">Already registered? <a href="/login">Proceed to login.</a></p>
    </div>
</form>

<style>
    * {
        padding: 0;
        margin: 0;
        border: 0;
    }
    :global(body){
        padding: 0;
        margin: 0;
        background: url('../../lib/assets/wallpaper.jpg');
    }
    p {
        padding: 7px 0 7px 0;
        font-weight: 500;
        font-size: 12pt;
    }
    h1 {
        font-weight:500;
        color: #888888;
        font-size:20pt;
        margin:7px 5px 8px 8px;
    }
    form.register{
        width:900px;
        margin: 150px auto 0px auto;
        height:470px;
        padding:5px;
        background: rgba(255, 255, 255, 0.2);
        backdrop-filter: blur(10px);
        -webkit-backdrop-filter: blur(10px);
        box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    }
    form p{
        font-size: 12pt;
        clear:both;
        margin: 0;
        color:gray;
        padding:4px;
    }
    form.register fieldset.row1
    {
        width:800px;
        padding:5px;
        float:left;
        /* border-top:1px solid #F5F5F5; */
        margin-bottom:15px;
    }
    form.register fieldset.row1 label{
        width:140px;
        float: left;
        text-align: right;
        margin-right: 6px;
        margin-top:2px;
    }
    form.register fieldset.row2
    {
        /* border-top:1px solid #F1F1F1;
        border-right:1px solid #F1F1F1; */
        height:115px;
        padding:5px;
        float:left;
    }
    form.register fieldset.row3
    {
        /* border-top:1px solid #F1F1F1; */
        padding:5px;
        float:left;
        margin-bottom:15px;
        width:450px;
    }
    form.register legend
    {
        color: #2563eb;
        padding:2px;
        margin-left: 14px;
        font-weight:bold;
        font-size: 20px;
        font-weight:100;
    }
    form.register label{
        color:#444;
        width:98px;
        float: left;
        text-align: right;
        margin-right: 6px;
        margin-top:2px;
    }
    form.register label.optional{
        float: left;
        text-align: right;
        margin-right: 6px;
        margin-top:2px;
        color: #A3A3A3;
    }
    form.register label.obinfo{
        float:right;
        padding:3px;
        font-style:italic;
    }
    form.register input{
        padding: 2px 10px;
        width: 140px;
        float: left;
        margin-right: 5px;
        background: rgba(255, 255, 255, 0.2);
        backdrop-filter: blur(10px);
        -webkit-backdrop-filter: blur(10px);
        box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    }
    form.register input.long{
        padding: 2px 10px;
        width: 247px;
        background: rgba(255, 255, 255, 0.2);
        backdrop-filter: blur(10px);
        -webkit-backdrop-filter: blur(10px);
        box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    }
    form.register input[type=radio]
    {
        float:left;
        width:15px;
    }
    form.register label.gender{
        margin-top:-1px;
        margin-bottom:2px;
        width:34px;
        float:left;
        text-align:left;
        line-height:19px;
    }
    form.register input[type=text]
    {
        border: 1px solid #E1E1E1;
        height: 18px;
    }
    .register-btn 
    {
        text-align: center;
        width: 100%;
        clear: both;
        padding-top: 40px;
    }
    .button
    {
        display: inline-block;
        padding: 8px 20px;
        text-decoration: none;
        color: #888888;
        cursor: pointer;
        font-size:18px;
        margin:10px auto;
        background: rgba(255, 255, 255, 0.2);
        backdrop-filter: blur(10px);
        -webkit-backdrop-filter: blur(10px);
        box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    }
    .button:hover {
        background-color: rgba(37, 99, 235, 0.7);
        color: #E1E1E1;
    }
    form.register input[type=text].year
    {
        border: 1px solid #E1E1E1;
        height: 15px;
        width:30px;
    }
    form.register select
    {
        width: 130px;
        padding-top: 2px;
        padding-bottom: 2px;
        float:left;
        margin-bottom:3px;
        margin-right:5px;
        background: rgba(255, 255, 255, 0.2);
        backdrop-filter: blur(10px);
        -webkit-backdrop-filter: blur(10px);
        box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    }
    form.register select.date
    {
        width: 40px;
        padding-top: 2px;
        padding-bottom: 2px;
        background: rgba(255, 255, 255, 0.2);
        backdrop-filter: blur(10px);
        -webkit-backdrop-filter: blur(10px);
        box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    }
    input:focus, select:focus{
        background-color: #efffe0;
    }
    .message {
        margin: 15px 0 0;
        color: #444;
        font-size: 12px;
    }
    .message a {
        color: #2563eb;
        text-decoration: none;
    }
</style>