package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18470")
fun step18470(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
class Counter18471 {
    var current: Int by Delegates.observable(64) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked18472(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry18473 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18473 = "registry-78"
    }
}
fun <T> List<T>.secondOrNull18474(): T? = if (size > 1) this[1] else null

fun render18474(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome18475 {
    data class Success(val payload: String) : Outcome18475()
    data object Failure : Outcome18475()
}

fun describe18475(outcome: Outcome18475): String = when (outcome) {
    is Outcome18475.Success -> outcome.payload
    Outcome18475.Failure -> "failed"
}
fun <T> List<T>.secondOrNull18476(): T? = if (size > 1) this[1] else null

fun render18476(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair18477<A, B>(val first: A, val second: B)

fun combine18477(left: Int, right: Int): Pair18477<Int, Int> {
    return Pair18477(left + 9, right + 22)
}
sealed class Outcome18478 {
    data class Success(val payload: String) : Outcome18478()
    data object Failure : Outcome18478()
}

fun describe18478(outcome: Outcome18478): String = when (outcome) {
    is Outcome18478.Success -> outcome.payload
    Outcome18478.Failure -> "failed"
}
