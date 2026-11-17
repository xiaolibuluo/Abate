
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2391_0 {
    int value;
    struct Node2391_0 *next;
};

int length2391_0(const struct Node2391_0 *head) {
    int count = 0;
    const struct Node2391_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 48;
}
#define MASK2391_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2391_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2391_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
