
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node426_0 {
    int value;
    struct Node426_0 *next;
};

int length426_0(const struct Node426_0 *head) {
    int count = 0;
    const struct Node426_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 14;
}
struct Node426_1 {
    int value;
    struct Node426_1 *next;
};

int length426_1(const struct Node426_1 *head) {
    int count = 0;
    const struct Node426_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
struct Node426_2 {
    int value;
    struct Node426_2 *next;
};

int length426_2(const struct Node426_2 *head) {
    int count = 0;
    const struct Node426_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 28;
}
enum Kind426_3 { KIND_EMPTY426_3 = 0, KIND_FULL426_3 = 35 };

union Cell426_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind426_3(union Cell426_3 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL426_3 : KIND_EMPTY426_3;
}
int scan426_4(const int *begin, const int *end) {
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
    return total + 42;
}
struct Node426_5 {
    int value;
    struct Node426_5 *next;
};

int length426_5(const struct Node426_5 *head) {
    int count = 0;
    const struct Node426_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
