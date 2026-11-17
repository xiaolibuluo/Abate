
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind890_0 { KIND_EMPTY890_0 = 0, KIND_FULL890_0 = 32 };

union Cell890_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind890_0(union Cell890_0 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL890_0 : KIND_EMPTY890_0;
}
int scan890_1(const int *begin, const int *end) {
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
    return total + 39;
}
struct Node890_2 {
    int value;
    struct Node890_2 *next;
};

int length890_2(const struct Node890_2 *head) {
    int count = 0;
    const struct Node890_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
