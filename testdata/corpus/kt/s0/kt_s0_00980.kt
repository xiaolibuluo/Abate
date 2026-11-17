package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8030 {
    data class Success(val payload: String) : Outcome8030()
    data object Failure : Outcome8030()
}

fun describe8030(outcome: Outcome8030): String = when (outcome) {
    is Outcome8030.Success -> outcome.payload
    Outcome8030.Failure -> "failed"
}
fun <T> List<T>.secondOrNull8031(): T? = if (size > 1) this[1] else null

fun render8031(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair8032<A, B>(val first: A, val second: B)

fun combine8032(left: Int, right: Int): Pair8032<Int, Int> {
    return Pair8032(left + 79, right + 40)
}
sealed class Outcome8033 {
    data class Success(val payload: String) : Outcome8033()
    data object Failure : Outcome8033()
}

fun describe8033(outcome: Outcome8033): String = when (outcome) {
    is Outcome8033.Success -> outcome.payload
    Outcome8033.Failure -> "failed"
}
sealed class Outcome8034 {
    data class Success(val payload: String) : Outcome8034()
    data object Failure : Outcome8034()
}

fun describe8034(outcome: Outcome8034): String = when (outcome) {
    is Outcome8034.Success -> outcome.payload
    Outcome8034.Failure -> "failed"
}
