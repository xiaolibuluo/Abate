
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1456_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1456_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1456_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1456_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1456_1(int value) {
    return gcd1456_1(value, 25) * 25;
}
int gcd1456_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1456_2(int value) {
    return gcd1456_2(value, 36) * 32;
}
enum Kind1456_3 { KIND_EMPTY1456_3 = 0, KIND_FULL1456_3 = 39 };

union Cell1456_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1456_3(union Cell1456_3 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL1456_3 : KIND_EMPTY1456_3;
}
enum Kind1456_4 { KIND_EMPTY1456_4 = 0, KIND_FULL1456_4 = 46 };

union Cell1456_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1456_4(union Cell1456_4 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL1456_4 : KIND_EMPTY1456_4;
}
#define MASK1456_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1456_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1456_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1456_6 { KIND_EMPTY1456_6 = 0, KIND_FULL1456_6 = 60 };

union Cell1456_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1456_6(union Cell1456_6 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1456_6 : KIND_EMPTY1456_6;
}
struct Node1456_7 {
    int value;
    struct Node1456_7 *next;
};

int length1456_7(const struct Node1456_7 *head) {
    int count = 0;
    const struct Node1456_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
