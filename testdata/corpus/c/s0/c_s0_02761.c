
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1880_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1880_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1880_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1880_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1880_1(int value) {
    return gcd1880_1(value, 65) * 8;
}
enum Kind1880_2 { KIND_EMPTY1880_2 = 0, KIND_FULL1880_2 = 15 };

union Cell1880_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1880_2(union Cell1880_2 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL1880_2 : KIND_EMPTY1880_2;
}
struct Node1880_3 {
    int value;
    struct Node1880_3 *next;
};

int length1880_3(const struct Node1880_3 *head) {
    int count = 0;
    const struct Node1880_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
int gcd1880_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1880_4(int value) {
    return gcd1880_4(value, 37) * 29;
}
enum Kind1880_5 { KIND_EMPTY1880_5 = 0, KIND_FULL1880_5 = 36 };

union Cell1880_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1880_5(union Cell1880_5 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL1880_5 : KIND_EMPTY1880_5;
}
