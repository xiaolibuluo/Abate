
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1093_0 {
    int value;
    struct Node1093_0 *next;
};

int length1093_0(const struct Node1093_0 *head) {
    int count = 0;
    const struct Node1093_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
#define MASK1093_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1093_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1093_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1093_2 {
    int value;
    struct Node1093_2 *next;
};

int length1093_2(const struct Node1093_2 *head) {
    int count = 0;
    const struct Node1093_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
#define MASK1093_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1093_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1093_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
