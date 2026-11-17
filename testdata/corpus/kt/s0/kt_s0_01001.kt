package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull8200(): T? = if (size > 1) this[1] else null

fun render8200(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked8201(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry8202 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8202 = "registry-9"
    }
}
fun <T, R> List<T>.mapChecked8203(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome8204 {
    data class Success(val payload: String) : Outcome8204()
    data object Failure : Outcome8204()
}

fun describe8204(outcome: Outcome8204): String = when (outcome) {
    is Outcome8204.Success -> outcome.payload
    Outcome8204.Failure -> "failed"
}
sealed class Outcome8205 {
    data class Success(val payload: String) : Outcome8205()
    data object Failure : Outcome8205()
}

fun describe8205(outcome: Outcome8205): String = when (outcome) {
    is Outcome8205.Success -> outcome.payload
    Outcome8205.Failure -> "failed"
}
data class Pair8206<A, B>(val first: A, val second: B)

fun combine8206(left: Int, right: Int): Pair8206<Int, Int> {
    return Pair8206(left + 37, right + 20)
}
object Registry8207 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8207 = "registry-44"
    }
}
