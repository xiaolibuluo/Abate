package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6800(): T? = if (size > 1) this[1] else null

fun render6800(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step6801")
fun step6801(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 25
}
class Counter6802 {
    var current: Int by Delegates.observable(32) { _, _, new -> println(new) }

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
data class Pair6803<A, B>(val first: A, val second: B)

fun combine6803(left: Int, right: Int): Pair6803<Int, Int> {
    return Pair6803(left + 39, right + 44)
}
sealed class Outcome6804 {
    data class Success(val payload: String) : Outcome6804()
    data object Failure : Outcome6804()
}

fun describe6804(outcome: Outcome6804): String = when (outcome) {
    is Outcome6804.Success -> outcome.payload
    Outcome6804.Failure -> "failed"
}
object Registry6805 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6805 = "registry-53"
    }
}
data class Pair6806<A, B>(val first: A, val second: B)

fun combine6806(left: Int, right: Int): Pair6806<Int, Int> {
    return Pair6806(left + 60, right + 16)
}
sealed class Outcome6807 {
    data class Success(val payload: String) : Outcome6807()
    data object Failure : Outcome6807()
}

fun describe6807(outcome: Outcome6807): String = when (outcome) {
    is Outcome6807.Success -> outcome.payload
    Outcome6807.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked6808(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
