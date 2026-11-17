
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1434_0 {
    int value;
    struct Node1434_0 *next;
};

int length1434_0(const struct Node1434_0 *head) {
    int count = 0;
    const struct Node1434_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
struct Node1434_1 {
    int value;
    struct Node1434_1 *next;
};

int length1434_1(const struct Node1434_1 *head) {
    int count = 0;
    const struct Node1434_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
struct Node1434_2 {
    int value;
    struct Node1434_2 *next;
};

int length1434_2(const struct Node1434_2 *head) {
    int count = 0;
    const struct Node1434_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
#define MASK1434_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1434_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1434_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
