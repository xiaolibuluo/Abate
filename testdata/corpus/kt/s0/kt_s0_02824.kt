package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23290<A, B>(val first: A, val second: B)

fun combine23290(left: Int, right: Int): Pair23290<Int, Int> {
    return Pair23290(left + 18, right + 25)
}
@JvmName("step23291")
fun step23291(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 25
}
object Registry23292 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23292 = "registry-32"
    }
}
fun <T, R> List<T>.mapChecked23293(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked23294(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome23295 {
    data class Success(val payload: String) : Outcome23295()
    data object Failure : Outcome23295()
}

fun describe23295(outcome: Outcome23295): String = when (outcome) {
    is Outcome23295.Success -> outcome.payload
    Outcome23295.Failure -> "failed"
}
fun <T> List<T>.secondOrNull23296(): T? = if (size > 1) this[1] else null

fun render23296(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter23297 {
    var current: Int by Delegates.observable(67) { _, _, new -> println(new) }

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
