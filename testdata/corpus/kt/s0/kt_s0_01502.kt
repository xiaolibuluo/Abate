package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry12430 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12430 = "registry-62"
    }
}
fun <T> List<T>.secondOrNull12431(): T? = if (size > 1) this[1] else null

fun render12431(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step12432")
fun step12432(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
class Counter12433 {
    var current: Int by Delegates.observable(83) { _, _, new -> println(new) }

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
sealed class Outcome12434 {
    data class Success(val payload: String) : Outcome12434()
    data object Failure : Outcome12434()
}

fun describe12434(outcome: Outcome12434): String = when (outcome) {
    is Outcome12434.Success -> outcome.payload
    Outcome12434.Failure -> "failed"
}
