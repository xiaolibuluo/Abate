package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked2410(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked2411(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter2412 {
    var current: Int by Delegates.observable(48) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
data class Pair2413<A, B>(val first: A, val second: B)

fun combine2413(left: Int, right: Int): Pair2413<Int, Int> {
    return Pair2413(left + 55, right + 21)
}
fun <T, R> List<T>.mapChecked2414(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair2415<A, B>(val first: A, val second: B)

fun combine2415(left: Int, right: Int): Pair2415<Int, Int> {
    return Pair2415(left + 69, right + 43)
}
object Registry2416 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2416 = "registry-76"
    }
}
fun <T> List<T>.secondOrNull2417(): T? = if (size > 1) this[1] else null

fun render2417(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome2418 {
    data class Success(val payload: String) : Outcome2418()
    data object Failure : Outcome2418()
}

fun describe2418(outcome: Outcome2418): String = when (outcome) {
    is Outcome2418.Success -> outcome.payload
    Outcome2418.Failure -> "failed"
}
