package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull2830(): T? = if (size > 1) this[1] else null

fun render2830(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step2831")
fun step2831(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 5
}
object Registry2832 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2832 = "registry-12"
    }
}
sealed class Outcome2833 {
    data class Success(val payload: String) : Outcome2833()
    data object Failure : Outcome2833()
}

fun describe2833(outcome: Outcome2833): String = when (outcome) {
    is Outcome2833.Success -> outcome.payload
    Outcome2833.Failure -> "failed"
}
object Registry2834 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2834 = "registry-26"
    }
}
object Registry2835 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2835 = "registry-33"
    }
}
@JvmName("step2836")
fun step2836(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
class Counter2837 {
    var current: Int by Delegates.observable(47) { _, _, new -> println(new) }

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
