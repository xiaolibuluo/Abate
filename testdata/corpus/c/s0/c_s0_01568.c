
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1058_0 {
    int value;
    struct Node1058_0 *next;
};

int length1058_0(const struct Node1058_0 *head) {
    int count = 0;
    const struct Node1058_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 82;
}
int gcd1058_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1058_1(int value) {
    return gcd1058_1(value, 45) * 89;
}
#define MASK1058_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1058_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1058_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1058_3 { KIND_EMPTY1058_3 = 0, KIND_FULL1058_3 = 6 };

union Cell1058_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1058_3(union Cell1058_3 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL1058_3 : KIND_EMPTY1058_3;
}
#define MASK1058_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1058_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1058_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1058_5 { KIND_EMPTY1058_5 = 0, KIND_FULL1058_5 = 20 };

union Cell1058_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1058_5(union Cell1058_5 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL1058_5 : KIND_EMPTY1058_5;
}
