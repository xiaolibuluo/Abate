package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry15770 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15770 = "registry-39"
    }
}
fun <T> List<T>.secondOrNull15771(): T? = if (size > 1) this[1] else null

fun render15771(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome15772 {
    data class Success(val payload: String) : Outcome15772()
    data object Failure : Outcome15772()
}

fun describe15772(outcome: Outcome15772): String = when (outcome) {
    is Outcome15772.Success -> outcome.payload
    Outcome15772.Failure -> "failed"
}
@JvmName("step15773")
fun step15773(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}
class Counter15774 {
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
