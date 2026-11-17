package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull5120(): T? = if (size > 1) this[1] else null

fun render5120(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked5121(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome5122 {
    data class Success(val payload: String) : Outcome5122()
    data object Failure : Outcome5122()
}

fun describe5122(outcome: Outcome5122): String = when (outcome) {
    is Outcome5122.Success -> outcome.payload
    Outcome5122.Failure -> "failed"
}
sealed class Outcome5123 {
    data class Success(val payload: String) : Outcome5123()
    data object Failure : Outcome5123()
}

fun describe5123(outcome: Outcome5123): String = when (outcome) {
    is Outcome5123.Success -> outcome.payload
    Outcome5123.Failure -> "failed"
}
data class Pair5124<A, B>(val first: A, val second: B)

fun combine5124(left: Int, right: Int): Pair5124<Int, Int> {
    return Pair5124(left + 93, right + 38)
}
