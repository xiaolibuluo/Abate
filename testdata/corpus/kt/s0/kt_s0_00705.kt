package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry5800 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5800 = "registry-76"
    }
}
fun <T> List<T>.secondOrNull5801(): T? = if (size > 1) this[1] else null

fun render5801(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step5802")
fun step5802(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
class Counter5803 {
    var current: Int by Delegates.observable(97) { _, _, new -> println(new) }

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
data class Pair5804<A, B>(val first: A, val second: B)

fun combine5804(left: Int, right: Int): Pair5804<Int, Int> {
    return Pair5804(left + 7, right + 26)
}
fun <T, R> List<T>.mapChecked5805(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome5806 {
    data class Success(val payload: String) : Outcome5806()
    data object Failure : Outcome5806()
}

fun describe5806(outcome: Outcome5806): String = when (outcome) {
    is Outcome5806.Success -> outcome.payload
    Outcome5806.Failure -> "failed"
}
data class Pair5807<A, B>(val first: A, val second: B)

fun combine5807(left: Int, right: Int): Pair5807<Int, Int> {
    return Pair5807(left + 28, right + 59)
}
fun <T, R> List<T>.mapChecked5808(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
