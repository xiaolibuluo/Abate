
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind376_0 { KIND_EMPTY376_0 = 0, KIND_FULL376_0 = 43 };

union Cell376_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind376_0(union Cell376_0 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL376_0 : KIND_EMPTY376_0;
}
struct Node376_1 {
    int value;
    struct Node376_1 *next;
};

int length376_1(const struct Node376_1 *head) {
    int count = 0;
    const struct Node376_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
enum Kind376_2 { KIND_EMPTY376_2 = 0, KIND_FULL376_2 = 57 };

union Cell376_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind376_2(union Cell376_2 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL376_2 : KIND_EMPTY376_2;
}
#define MASK376_3(v) (((v) << 4) | ((v) & 0x0F))

int classify376_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK376_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant376_4(int value, ...) {
    return value + 71;
}
#else
int variant376_4(int value) {
    return value + 71;
}
#endif
struct Node376_5 {
    int value;
    struct Node376_5 *next;
};

int length376_5(const struct Node376_5 *head) {
    int count = 0;
    const struct Node376_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 78;
}
