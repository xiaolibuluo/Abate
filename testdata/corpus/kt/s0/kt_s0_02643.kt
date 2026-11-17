package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked21840(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull21841(): T? = if (size > 1) this[1] else null

fun render21841(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked21842(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair21843<A, B>(val first: A, val second: B)

fun combine21843(left: Int, right: Int): Pair21843<Int, Int> {
    return Pair21843(left + 94, right + 19)
}
object Registry21844 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21844 = "registry-101"
    }
}
fun <T, R> List<T>.mapChecked21845(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome21846 {
    data class Success(val payload: String) : Outcome21846()
    data object Failure : Outcome21846()
}

fun describe21846(outcome: Outcome21846): String = when (outcome) {
    is Outcome21846.Success -> outcome.payload
    Outcome21846.Failure -> "failed"
}
class Counter21847 {
    var current: Int by Delegates.observable(25) { _, _, new -> println(new) }

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
