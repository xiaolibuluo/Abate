
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd526_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale526_0(int value) {
    return gcd526_0(value, 31) * 53;
}
#define MASK526_1(v) (((v) << 4) | ((v) & 0x0F))

int classify526_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK526_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind526_2 { KIND_EMPTY526_2 = 0, KIND_FULL526_2 = 67 };

union Cell526_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind526_2(union Cell526_2 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL526_2 : KIND_EMPTY526_2;
}
struct Node526_3 {
    int value;
    struct Node526_3 *next;
};

int length526_3(const struct Node526_3 *head) {
    int count = 0;
    const struct Node526_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 74;
}
#define MASK526_4(v) (((v) << 4) | ((v) & 0x0F))

int classify526_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK526_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node526_5 {
    int value;
    struct Node526_5 *next;
};

int length526_5(const struct Node526_5 *head) {
    int count = 0;
    const struct Node526_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
int gcd526_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale526_6(int value) {
    return gcd526_6(value, 36) * 95;
}
