
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1940_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1940_0(int value) {
    return gcd1940_0(value, 47) * 5;
}
int gcd1940_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1940_1(int value) {
    return gcd1940_1(value, 58) * 12;
}
struct Node1940_2 {
    int value;
    struct Node1940_2 *next;
};

int length1940_2(const struct Node1940_2 *head) {
    int count = 0;
    const struct Node1940_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}
enum Kind1940_3 { KIND_EMPTY1940_3 = 0, KIND_FULL1940_3 = 26 };

union Cell1940_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1940_3(union Cell1940_3 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1940_3 : KIND_EMPTY1940_3;
}
struct Node1940_4 {
    int value;
    struct Node1940_4 *next;
};

int length1940_4(const struct Node1940_4 *head) {
    int count = 0;
    const struct Node1940_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
enum Kind1940_5 { KIND_EMPTY1940_5 = 0, KIND_FULL1940_5 = 40 };

union Cell1940_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1940_5(union Cell1940_5 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL1940_5 : KIND_EMPTY1940_5;
}
#define MASK1940_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1940_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1940_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
