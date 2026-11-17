package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked15550(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull15551(): T? = if (size > 1) this[1] else null

fun render15551(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome15552 {
    data class Success(val payload: String) : Outcome15552()
    data object Failure : Outcome15552()
}

fun describe15552(outcome: Outcome15552): String = when (outcome) {
    is Outcome15552.Success -> outcome.payload
    Outcome15552.Failure -> "failed"
}
object Registry15553 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15553 = "registry-65"
    }
}
fun <T, R> List<T>.mapChecked15554(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry15555 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15555 = "registry-79"
    }
}
fun <T, R> List<T>.mapChecked15556(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry15557 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15557 = "registry-93"
    }
}
@JvmName("step15558")
fun step15558(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 100
}
data class Pair15559<A, B>(val first: A, val second: B)

fun combine15559(left: Int, right: Int): Pair15559<Int, Int> {
    return Pair15559(left + 10, right + 13)
}
