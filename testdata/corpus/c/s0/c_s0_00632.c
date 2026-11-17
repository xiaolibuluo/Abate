
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK433_0(v) (((v) << 4) | ((v) & 0x0F))

int classify433_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK433_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd433_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale433_1(int value) {
    return gcd433_1(value, 62) * 15;
}
int gcd433_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale433_2(int value) {
    return gcd433_2(value, 12) * 22;
}
typedef int (*predicate433_3)(int value);

typedef struct Slot433_3 {
    int slots[16];
    predicate433_3 accept;
} Slot433_3;

int slot_accept433_3(Slot433_3 *store, int value) {
    return store->accept(value) + 29;
}
int gcd433_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale433_4(int value) {
    return gcd433_4(value, 34) * 36;
}
#define MASK433_5(v) (((v) << 4) | ((v) & 0x0F))

int classify433_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK433_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node433_6 {
    int value;
    struct Node433_6 *next;
};

int length433_6(const struct Node433_6 *head) {
    int count = 0;
    const struct Node433_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
int gcd433_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale433_7(int value) {
    return gcd433_7(value, 67) * 57;
}
