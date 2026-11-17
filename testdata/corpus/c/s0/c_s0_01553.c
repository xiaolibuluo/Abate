
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1049_0 {
    int value;
    struct Node1049_0 *next;
};

int length1049_0(const struct Node1049_0 *head) {
    int count = 0;
    const struct Node1049_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}
struct Node1049_1 {
    int value;
    struct Node1049_1 *next;
};

int length1049_1(const struct Node1049_1 *head) {
    int count = 0;
    const struct Node1049_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
enum Kind1049_2 { KIND_EMPTY1049_2 = 0, KIND_FULL1049_2 = 76 };

union Cell1049_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1049_2(union Cell1049_2 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL1049_2 : KIND_EMPTY1049_2;
}
int scan1049_3(const int *begin, const int *end) {
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
    return total + 83;
}
