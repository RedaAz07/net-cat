# 📡 Network (Réseau Informatique)

## 🖥️ C'est quoi un Network ?
Un **network (réseau informatique)** est un ensemble de plusieurs **appareils connectés** entre eux, soit par **câble (wired connection)**, soit sans fil **(wireless connection - Wi-Fi)**, afin de communiquer et partager des ressources.

---

## ❓ Pourquoi avons-nous besoin d'un Network ?
```
✅ Partager du matériel (ex: une imprimante connectée à plusieurs PC).
✅ Partager des données (ex: fichiers, projets en entreprise).
✅ Protéger les données (ex: backup = sauvegarde des fichiers sur un disque externe, serveur distant ou cloud).
```

---

## 🔌 Network Media (Média de Transmission)
Pour qu'au **moins deux PC** puissent être connectés et former un réseau, il faut un **moyen de connexion physique** :
```
📡 Soit une connexion par câble (wired).
📶 Soit une connexion sans fil (wireless - Wi-Fi).
```

---

## 🔗 Types de Câbles

### 1️⃣ **Copper Cables (Câbles en Cuivre)**
📌 Il en existe deux types :
- **Coaxial Cable** : Anciennement utilisé pour la TV et Internet (rarement utilisé aujourd'hui).
- **Twisted Pair Cable** : C'est le plus courant, il relie les appareils aux routeurs/switchs (ex: Ethernet RJ45).
- 🏠 **Distance maximale** : Environ **100m** avant perte de signal.

### 2️⃣ **Fiber Optic Cables (Fibre Optique)**
📌 Utilisé pour des connexions ultra-rapides, fonctionne en envoyant des **impulsions lumineuses** à travers des fils en verre.

```
✅ La carte réseau (NIC - Network Interface Card) adapte le signal selon le type de réseau :
   - 🌐 Fibre optique 🔦 → Convertit les données en signaux lumineux.
   - ⚡ Ethernet (câble RJ45) → Convertit les données en signaux électriques.
   - 📡 Wi-Fi → Convertit les données en ondes radio.
```

---

## 🔄 **Les Types de Topologie Réseau**

### 1️⃣ **Bus (الباص) 🚌**
📌 الأجهزة كاتكون مربوطة فخط واحد وكيتشاركو نفس الكابل.
✅ ساهل فالتنفيذ.
❌ إيلا وقع مشكل فالكابل، الشبكة كاملة تطيح.

### 2️⃣ **Star (النجمة) ⭐**
📌 كل الأجهزة مربوطة بسيرفر أو سويتش فالنص.
✅ مستقرة، إيلا خسر جهاز واحد، الشبكة كاتبقى خدامة.
❌ إيلا طاح السيرفر/السويتش، الشبكة كاملة تطيح.

### 3️⃣ **Ring (الحلقية) 🔄**
📌 الأجهزة مربوطة فشكل دائرة، والبيانات كاتمشي فاتجاه معين.
✅ مزيانة للشبكات لي كاتحتاج تنظيم البيانات.
❌ إيلا خسر جهاز واحد، كاين احتمال توقف الشبكة.

### 4️⃣ **Mesh (الشبكية) 🕸️**
📌 كل جهاز كايتواصل مع جميع الأجهزة مباشرة.
✅ قوية بزاف، ماكايناش نقطة ضعف.
❌ غالية ومعقدة فالتنفيذ.

---

## Network Types
```
1. LAN (Local Area Network) - الشبكة المحلية
شبكة صغيرة كاتربط الأجهزة فبلاصا محدودة بحال الدار ولا الشركة.
🖥 مثال: الحواسيب فمكتب مربوطين بكابل إيثرنت.

2. WAN (Wide Area Network) - الشبكة الواسعة
كاتربط بزاف ديال الشبكات فمناطق بعيدة باستعمال الإنترنت ولا خطوط اتصال خاصة.
🌍 مثال: الإنترنت، الشركات اللي عندهم فروع متواصلين بـ VPN.

3. MAN (Metropolitan Area Network) - شبكة المدينة
شبكة وسطانية، كاتكون أكبر من LAN وصغيرة على WAN، كاتخدم فمدينة ولا فجامعة.
🏙 مثال: Wi-Fi ديال شي مدينة ولا شبكة جامعة كبيرة.

4. PAN (Personal Area Network) - الشبكة الشخصية
كاتربط الأجهزة اللي قريبة لبعضها شي مترات، غالباً عبر Bluetooth ولا USB.
📱 مثال: تليفون مربوط بسماعات بلوتوث.

5. WLAN (Wireless LAN) - شبكة محلية لاسلكية
LAN بلا كابلات، كاتخدم بالـ Wi-Fi.
📡 مثال: Wi-Fi فالقهوة.

6. SAN (Storage Area Network) - شبكة التخزين
شبكة سريعة مخصصة للتخزين والسيرفرات.
💾 مثال: Cloud storage ديال الشركات الكبيرة.

7. VPN (Virtual Private Network) - الشبكة الافتراضية الخاصة
كاتسمح لك تدخل شي شبكة بعيدة بشكل آمن ومشفر.
🔒 مثال: موظف خدام من الدار وكيستعمل VPN باش يدخل لسيرفر الشركة.
```

---

## 📚 **Les 7 Couches du Modèle OSI**
```
1️⃣ **Physical (الفيزيائية)**: كاتحدد الكابلات، الإشارات، والتوصيلات المادية.
2️⃣ **Data Link (رابط البيانات)**: مسؤولة على إرسال البيانات بين الأجهزة مباشرة.
3️⃣ **Network (الشبكة)**: كاتحدد كيفاش البيانات كاتسافر بين الشبكات (ex: IP addresses).
4️⃣ **Transport (النقل)**: كاتضمن نقل البيانات بطريقة مضمونة (ex: TCP, UDP).
5️⃣ **Session (الجلسة)**: كاتدير التحكم فالجلسات بين المستخدمين.
6️⃣ **Presentation (التقديم)**: كاتحول البيانات باش تفهمها التطبيقات (ex: التشفير، الضغط).
7️⃣ **Application (التطبيقات)**: كاتتعامل مع التطبيقات اللي كايستعملها المستخدم (ex: HTTP, FTP).
```

---

## 🌐 **Le Modèle TCP/IP**
```
1️⃣ **Network Access (الوصول للشبكة)**: كاتحدد كيفاش الأجهزة كاتتواصل مع الشبكة مباشرة.
2️⃣ **Internet (الإنترنت)**: كاتهتم بتوجيه البيانات بين الشبكات (ex: IP, ICMP).
3️⃣ **Transport (النقل)**: مسؤولة على نقل البيانات بطريقة مضمونة أو سريعة (ex: TCP, UDP).
4️⃣ **Application (التطبيقات)**: كاتحدد البروتوكولات اللي كاتخدم بها التطبيقات (ex: HTTP, FTP, SMTP).
```

