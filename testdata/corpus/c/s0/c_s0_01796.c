
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1205_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1205_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1205_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1205_1 { KIND_EMPTY1205_1 = 0, KIND_FULL1205_1 = 60 };

union Cell1205_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1205_1(union Cell1205_1 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL1205_1 : KIND_EMPTY1205_1;
}
struct Node1205_2 {
    int value;
    struct Node1205_2 *next;
};

int length1205_2(const struct Node1205_2 *head) {
    int count = 0;
    const struct Node1205_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
int gcd1205_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1205_3(int value) {
    return gcd1205_3(value, 59) * 74;
}
int gcd1205_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1205_4(int value) {
    return gcd1205_4(value, 9) * 81;
}
struct Node1205_5 {
    int value;
    struct Node1205_5 *next;
};

int length1205_5(const struct Node1205_5 *head) {
    int count = 0;
    const struct Node1205_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
