package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair5080<A, B>(val first: A, val second: B)

fun combine5080(left: Int, right: Int): Pair5080<Int, Int> {
    return Pair5080(left + 13, right + 27)
}
fun <T, R> List<T>.mapChecked5081(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome5082 {
    data class Success(val payload: String) : Outcome5082()
    data object Failure : Outcome5082()
}

fun describe5082(outcome: Outcome5082): String = when (outcome) {
    is Outcome5082.Success -> outcome.payload
    Outcome5082.Failure -> "failed"
}
data class Pair5083<A, B>(val first: A, val second: B)

fun combine5083(left: Int, right: Int): Pair5083<Int, Int> {
    return Pair5083(left + 34, right + 60)
}
fun <T> List<T>.secondOrNull5084(): T? = if (size > 1) this[1] else null

fun render5084(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
