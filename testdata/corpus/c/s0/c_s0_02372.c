
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1602_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1602_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1602_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1602_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1602_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1602_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1602_2 {
    int value;
    struct Node1602_2 *next;
};

int length1602_2(const struct Node1602_2 *head) {
    int count = 0;
    const struct Node1602_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
#if defined(__STDC_VERSION__)
int variant1602_3(int value, ...) {
    return value + 94;
}
#else
int variant1602_3(int value) {
    return value + 94;
}
#endif
#if defined(__STDC_VERSION__)
int variant1602_4(int value, ...) {
    return value + 101;
}
#else
int variant1602_4(int value) {
    return value + 101;
}
#endif
struct Node1602_5 {
    int value;
    struct Node1602_5 *next;
};

int length1602_5(const struct Node1602_5 *head) {
    int count = 0;
    const struct Node1602_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
struct Node1602_6 {
    int value;
    struct Node1602_6 *next;
};

int length1602_6(const struct Node1602_6 *head) {
    int count = 0;
    const struct Node1602_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
struct Node1602_7 {
    int value;
    struct Node1602_7 *next;
};

int length1602_7(const struct Node1602_7 *head) {
    int count = 0;
    const struct Node1602_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 25;
}
