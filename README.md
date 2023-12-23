# CloudStart

## Defaults
Default configurations can be saved in the `$HOME/.config/cloudstart/defaults.yml` file

For detailed documentation visit [https://cloudinit.readthedocs.io/en/latest/reference/index.html](https://cloudinit.readthedocs.io/en/latest/reference/index.html)

**Caution:** There is a separate rule for [`passwd`](#userspasswd)

* [file](#file)
* * [file.path](#filepath)
* * [file.name](#filename)

* [packages](#packages)
* * [update](#packagesupdate)
* * [upgrade](#packagesupgrade)
* * [packages](#packagespackages)

* [ssh](#ssh)
* * [pwauth](#sshpwauth)
* * [disable_root](#sshdisable_root)
* * [no_fingerprints](#sshno_fingerprints)

* [user_options](#user_options)
* * [expiredate](#usersexpiredate)
* * [groups](#user_optionsgroups)
* * [lock_passwd](#userslock_passwd)
* * [inactive](#usersinactive)
* * [passwd](#userspasswd)
* * [no_create_home](#usersno_create_home)
* * [no_user_group](#usersno_user_group)
* * [no_log_init](#usersno_log_init)
* * [ssh_import_id](#usersssh_import_id)
* * [ssh_authorized_keys](#usersssh_authorized_keys)
* * [ssh_redirect_user](#usersssh_redirect_user)
* * [sudo](#userssudo)

* [users](#users)
* * [name](#usersname)
* * [expiredate](#expiredate)
* * [gecos](#usersgecos)
* * [homedir](#usershomedir)
* * [primary_group](#usersprimary_group)
* * [groups](#usersgroups)
* * [selinux_user](#usersselinux_user)
* * [lock_passwd](#userslock_passwd)
* * [inactive](#usersinactive)
* * [passwd](#userspasswd)
* * [no_create_home](#usersno_create_home)
* * [no_user_group](#usersno_user_group)
* * [no_log_init](#usersno_log_init)
* * [ssh_import_id](#usersssh_import_id)
* * [ssh_authorized_keys](#usersssh_authorized_keys)
* * [ssh_redirect_user](#usersssh_redirect_user)
* * [sudo](#userssudo)
* * [system](#userssystem)
* * [snapuser](#userssnapuser)

* [runcmds](#runcmds)
* * [name](#runcmdsname)
* * [cmds](#runcmdscmds)

---

### `file`

Sets the default path and name of the output config file

At the first start, the path and the file are created if it does not exist

Automatically created file:
```yml
file:
  path: .
  name: cloud-config
```

<br/>

### `file.path`
* Typ: `string`
* Default: `.`

Sets the path

<br/>

### `file.name`
* Typ: `string`
* Default: `cloud-init`

Sets the (name without the `.yml` extension)

---

### `packages`
Sets update, upgrade and standard packages to be installed

Example:
```yml
packages:
  update: true
  upgrade: true
  packages:
    - git
    - vim
```

<br/>

### `packages.update`
* Typ: `boolean`

Update apt database on first boot (run 'apt-get update')

<br/>

### `packages.upgrade`
* Typ: `boolean`

Upgrade the instance on first boot

<br/>

### `packages.packages`
* Typ: `list`

Install packages

---

### `ssh`

Sets the default values for ssh

Example:
```yml
ssh:
  pwauth: false
  disable_root: true
  no_fingerprints: false
```

<br/>

### `ssh.pwauth`
* Typ: `boolean`

Allow password authentication

<br/>

### `ssh.disable_root`
* Typ: `boolean`

If this is set, 'root' will not be able to ssh in and they will get a message to login instead as the default $user

<br/>

### `ssh.no_fingerprints`
* Typ: `boolean`

By default, the fingerprints of the authorized keys for the users cloud-init adds are printed to the console

Setting `no_ssh_fingerprints` to true suppresses this output

---

### `user_options`

Sets default values for the creation of users can simply be confirmed with enter

The documentation for the supported keys is the same as for defined users, see [`users`](#users)

Example:
```yml
user_options:
  expiredate: "2032-09-01"
  groups:
    - users
    - admin
  ssh_import_id:
    - lp:falcojr
    - gh:TheRealFalcon
  sudo:
    - ALL=(ALL) NOPASSWD:ALL
    - ALL=(ALL) NOPASSWD:/bin/mysql
```

<br/>

### `user_options.groups`
* Typ: `list`

Additional groups to add the user to

---

### `users`
* Typ: `list`

Sets predefined users that can be easily added with enter

Example:
```yml
users:
  - name: foobar
    expiredate: '2032-09-01'
    gecos: Foo B. Bar
    primary_group: foobar
    lock_passwd: false
    passwd: random
    ssh_authorized_keys:
      - ssh-rsa AAAA... csmith@fringe
  - name: cloudy
    gecos: Magic Cloud App Daemon User
    inactive: '5'
    system: true
```

### `users.name`
* **REQUIRED**
* Typ: `string`

The user's login name

<br/>

### `users.expiredate`
* Typ: `string`

Set the account expiration date

<br/>

### `users.gecos`
* Typ: `string`

The user name's real name, i.e. "Bob B. Smith"

<br/>

### `users.homedir`
* Typ: `string`

Set to the local path you want to use. Defaults to `/home/<username>`

<br/>

### `users.primary_group`
* Typ: `string`

Defaults to a new group created named after the user

<br/>

### `users.groups`
* Typ: `list`

Additional groups to add the user to

<br/>

### `users.selinux_user`
* Typ: `string`

The SELinux user for the user's login, such as "staff_u"

When this is omitted the system will select the default SELinux user

<br/>

### `users.lock_passwd`
* Typ: `boolean`

Lock the password to disable password login

<br/>

### `users.inactive`
* Typ: `string`

Number of days after password expires until account is disabled

<br/>

### `users.passwd`
* Typ: `string`

Set to `random` (creates a random 16 char password) or `assign` (password must be set)

The passwords are hashed with SHA-512 Crypt

All users are written with their password (in plain text) in the head of the config file **DON'T FORGET TO DELETE!**

<br/>

### `users.no_create_home`
* Typ: `boolean`

When set to true, do not create home directory

<br/>

### `users.no_user_group`
* Typ: `boolean`

When set to true, do not create a group named after the user

<br/>

### `users.no_log_init`
* Typ: `boolean`

When set to true, do not initialize lastlog and faillog database

<br/>

### `users.ssh_import_id`
* Typ: `list`

Import SSH ids

<br/>

### `users.ssh_authorized_keys`
* Typ: `list`

SSH authorized keys

<br/>

### `users.ssh_redirect_user`
* Typ: `boolean`

Set true to block ssh logins for cloud ssh public keys and emit a message redirecting logins to use `<default_username>` instead

This option only disables cloud provided public-keys

An error will be raised if ssh_authorized_keys or ssh_import_id is provided for the same user

<br/>

### `users.sudo`
* Typ: `list`

Sudo rule strings

<br/>

### `users.system`
* Typ: `boolean`

Create the user as a system user. This means no home directory.

<br/>

### `users.snapuser`
* Typ: `string`

Create a Snappy (Ubuntu-Core) user via the snap create-user ommand available on Ubuntu systems

If the user has an account on the Ubuntu SSO, specifying the email will allow snap to request a username and any public ssh keys and will import these into the system with username specified by SSO account

If `username` is not set in SSO, then username will be the shortname before the email domain

---

### `runcmds`
* Typ: `list`

Sets predefined commands that can be easily added with enter

Example:
```yml
runcmds:
  - name: 1G Swap
    cmds:
      - fallocate -l 1G /swapfile
      - chmod 600 /swapfile
      - ...
  - name: UFW defaults
    cmds:
      - ufw default deny incoming
      - ufw default allow outgoing
```

<br/>

### `runcmds.name`
* **REQUIRED**
* Typ: `string`

The name of the command block

<br/>

### `runcmds.cmds`
* **REQUIRED**
* Typ: `list`

The commands to run