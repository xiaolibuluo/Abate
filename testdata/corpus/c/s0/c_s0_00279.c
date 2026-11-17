
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node194_0 {
    int value;
    struct Node194_0 *next;
};

int length194_0(const struct Node194_0 *head) {
    int count = 0;
    const struct Node194_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
int scan194_1(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 12;
}
#define MASK194_2(v) (((v) << 4) | ((v) & 0x0F))

int classify194_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK194_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node194_3 {
    int value;
    struct Node194_3 *next;
};

int length194_3(const struct Node194_3 *head) {
    int count = 0;
    const struct Node194_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 26;
}
enum Kind194_4 { KIND_EMPTY194_4 = 0, KIND_FULL194_4 = 33 };

union Cell194_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind194_4(union Cell194_4 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL194_4 : KIND_EMPTY194_4;
}
enum Kind194_5 { KIND_EMPTY194_5 = 0, KIND_FULL194_5 = 40 };

union Cell194_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind194_5(union Cell194_5 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL194_5 : KIND_EMPTY194_5;
}
struct Node194_6 {
    int value;
    struct Node194_6 *next;
};

int length194_6(const struct Node194_6 *head) {
    int count = 0;
    const struct Node194_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
