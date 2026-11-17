package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22150<A, B>(val first: A, val second: B)

fun combine22150(left: Int, right: Int): Pair22150<Int, Int> {
    return Pair22150(left + 88, right + 20)
}
sealed class Outcome22151 {
    data class Success(val payload: String) : Outcome22151()
    data object Failure : Outcome22151()
}

fun describe22151(outcome: Outcome22151): String = when (outcome) {
    is Outcome22151.Success -> outcome.payload
    Outcome22151.Failure -> "failed"
}
fun <T> List<T>.secondOrNull22152(): T? = if (size > 1) this[1] else null

fun render22152(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked22153(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked22154(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step22155")
fun step22155(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
