
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK598_0(v) (((v) << 4) | ((v) & 0x0F))

int classify598_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK598_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK598_1(v) (((v) << 4) | ((v) & 0x0F))

int classify598_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK598_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node598_2 {
    int value;
    struct Node598_2 *next;
};

int length598_2(const struct Node598_2 *head) {
    int count = 0;
    const struct Node598_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
