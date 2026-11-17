
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1069_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1069_0(int value) {
    return gcd1069_0(value, 50) * 31;
}
enum Kind1069_1 { KIND_EMPTY1069_1 = 0, KIND_FULL1069_1 = 38 };

union Cell1069_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1069_1(union Cell1069_1 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL1069_1 : KIND_EMPTY1069_1;
}
struct Node1069_2 {
    int value;
    struct Node1069_2 *next;
};

int length1069_2(const struct Node1069_2 *head) {
    int count = 0;
    const struct Node1069_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
enum Kind1069_3 { KIND_EMPTY1069_3 = 0, KIND_FULL1069_3 = 52 };

union Cell1069_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1069_3(union Cell1069_3 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1069_3 : KIND_EMPTY1069_3;
}
#define MASK1069_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1069_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1069_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1069_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1069_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1069_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1069_6 { KIND_EMPTY1069_6 = 0, KIND_FULL1069_6 = 73 };

union Cell1069_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1069_6(union Cell1069_6 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL1069_6 : KIND_EMPTY1069_6;
}
enum Kind1069_7 { KIND_EMPTY1069_7 = 0, KIND_FULL1069_7 = 80 };

union Cell1069_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1069_7(union Cell1069_7 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL1069_7 : KIND_EMPTY1069_7;
}
