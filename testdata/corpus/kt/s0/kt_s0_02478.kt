package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair20390<A, B>(val first: A, val second: B)

fun combine20390(left: Int, right: Int): Pair20390<Int, Int> {
    return Pair20390(left + 31, right + 69)
}
class Counter20391 {
    var current: Int by Delegates.observable(38) { _, _, new -> println(new) }

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
@JvmName("step20392")
fun step20392(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
object Registry20393 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20393 = "registry-52"
    }
}
fun <T> List<T>.secondOrNull20394(): T? = if (size > 1) this[1] else null

fun render20394(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair20395<A, B>(val first: A, val second: B)

fun combine20395(left: Int, right: Int): Pair20395<Int, Int> {
    return Pair20395(left + 66, right + 63)
}
sealed class Outcome20396 {
    data class Success(val payload: String) : Outcome20396()
    data object Failure : Outcome20396()
}

fun describe20396(outcome: Outcome20396): String = when (outcome) {
    is Outcome20396.Success -> outcome.payload
    Outcome20396.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked20397(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked20398(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
