
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1078_0 {
    int value;
    struct Node1078_0 *next;
};

int length1078_0(const struct Node1078_0 *head) {
    int count = 0;
    const struct Node1078_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 51;
}
#define MASK1078_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1078_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1078_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
