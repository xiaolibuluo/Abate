
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node435_0 {
    int value;
    struct Node435_0 *next;
};

int length435_0(const struct Node435_0 *head) {
    int count = 0;
    const struct Node435_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
#define MASK435_1(v) (((v) << 4) | ((v) & 0x0F))

int classify435_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK435_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind435_2 { KIND_EMPTY435_2 = 0, KIND_FULL435_2 = 48 };

union Cell435_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind435_2(union Cell435_2 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL435_2 : KIND_EMPTY435_2;
}
#if defined(__STDC_VERSION__)
int variant435_3(int value, ...) {
    return value + 55;
}
#else
int variant435_3(int value) {
    return value + 55;
}
#endif
struct Node435_4 {
    int value;
    struct Node435_4 *next;
};

int length435_4(const struct Node435_4 *head) {
    int count = 0;
    const struct Node435_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}
#if defined(__STDC_VERSION__)
int variant435_5(int value, ...) {
    return value + 69;
}
#else
int variant435_5(int value) {
    return value + 69;
}
#endif
int gcd435_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale435_6(int value) {
    return gcd435_6(value, 9) * 76;
}
struct Node435_7 {
    int value;
    struct Node435_7 *next;
};

int length435_7(const struct Node435_7 *head) {
    int count = 0;
    const struct Node435_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 83;
}
