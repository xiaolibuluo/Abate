
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1068_0 { KIND_EMPTY1068_0 = 0, KIND_FULL1068_0 = 18 };

union Cell1068_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1068_0(union Cell1068_0 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL1068_0 : KIND_EMPTY1068_0;
}
struct Node1068_1 {
    int value;
    struct Node1068_1 *next;
};

int length1068_1(const struct Node1068_1 *head) {
    int count = 0;
    const struct Node1068_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 25;
}
#define MASK1068_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1068_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1068_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1068_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1068_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1068_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1068_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1068_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1068_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1068_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1068_5(int value) {
    return gcd1068_5(value, 37) * 53;
}
int gcd1068_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1068_6(int value) {
    return gcd1068_6(value, 48) * 60;
}
#define MASK1068_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1068_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1068_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
