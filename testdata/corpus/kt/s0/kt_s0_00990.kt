package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull8120(): T? = if (size > 1) this[1] else null

fun render8120(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter8121 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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
sealed class Outcome8122 {
    data class Success(val payload: String) : Outcome8122()
    data object Failure : Outcome8122()
}

fun describe8122(outcome: Outcome8122): String = when (outcome) {
    is Outcome8122.Success -> outcome.payload
    Outcome8122.Failure -> "failed"
}
@JvmName("step8123")
fun step8123(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 9
}
fun <T> List<T>.secondOrNull8124(): T? = if (size > 1) this[1] else null

fun render8124(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry8125 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8125 = "registry-23"
    }
}
