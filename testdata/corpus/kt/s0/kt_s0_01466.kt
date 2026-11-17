package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull12120(): T? = if (size > 1) this[1] else null

fun render12120(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome12121 {
    data class Success(val payload: String) : Outcome12121()
    data object Failure : Outcome12121()
}

fun describe12121(outcome: Outcome12121): String = when (outcome) {
    is Outcome12121.Success -> outcome.payload
    Outcome12121.Failure -> "failed"
}
@JvmName("step12122")
fun step12122(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
class Counter12123 {
    var current: Int by Delegates.observable(68) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked12124(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step12125")
fun step12125(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
object Registry12126 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12126 = "registry-89"
    }
}
data class Pair12127<A, B>(val first: A, val second: B)

fun combine12127(left: Int, right: Int): Pair12127<Int, Int> {
    return Pair12127(left + 96, right + 30)
}
@JvmName("step12128")
fun step12128(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 6
}
