package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step15530")
fun step15530(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
@JvmName("step15531")
fun step15531(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 25
}
object Registry15532 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15532 = "registry-32"
    }
}
object Registry15533 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15533 = "registry-39"
    }
}
class Counter15534 {
    var current: Int by Delegates.observable(46) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull15535(): T? = if (size > 1) this[1] else null

fun render15535(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair15536<A, B>(val first: A, val second: B)

fun combine15536(left: Int, right: Int): Pair15536<Int, Int> {
    return Pair15536(left + 60, right + 27)
}
sealed class Outcome15537 {
    data class Success(val payload: String) : Outcome15537()
    data object Failure : Outcome15537()
}

fun describe15537(outcome: Outcome15537): String = when (outcome) {
    is Outcome15537.Success -> outcome.payload
    Outcome15537.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked15538(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
