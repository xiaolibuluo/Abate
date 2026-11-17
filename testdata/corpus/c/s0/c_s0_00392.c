
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate265_0)(int value);

typedef struct Slot265_0 {
    int slots[16];
    predicate265_0 accept;
} Slot265_0;

int slot_accept265_0(Slot265_0 *store, int value) {
    return store->accept(value) + 55;
}
#define MASK265_1(v) (((v) << 4) | ((v) & 0x0F))

int classify265_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK265_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd265_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale265_2(int value) {
    return gcd265_2(value, 56) * 69;
}
int gcd265_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale265_3(int value) {
    return gcd265_3(value, 67) * 76;
}
struct Node265_4 {
    int value;
    struct Node265_4 *next;
};

int length265_4(const struct Node265_4 *head) {
    int count = 0;
    const struct Node265_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 83;
}
