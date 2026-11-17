package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry6260 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6260 = "registry-92"
    }
}
fun <T, R> List<T>.mapChecked6261(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair6262<A, B>(val first: A, val second: B)

fun combine6262(left: Int, right: Int): Pair6262<Int, Int> {
    return Pair6262(left + 9, right + 21)
}
fun <T, R> List<T>.mapChecked6263(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked6264(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome6265 {
    data class Success(val payload: String) : Outcome6265()
    data object Failure : Outcome6265()
}

fun describe6265(outcome: Outcome6265): String = when (outcome) {
    is Outcome6265.Success -> outcome.payload
    Outcome6265.Failure -> "failed"
}
fun <T> List<T>.secondOrNull6266(): T? = if (size > 1) this[1] else null

fun render6266(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry6267 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6267 = "registry-44"
    }
}
fun <T, R> List<T>.mapChecked6268(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
