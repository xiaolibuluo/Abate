package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked18330(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull18331(): T? = if (size > 1) this[1] else null

fun render18331(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry18332 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18332 = "registry-83"
    }
}
@JvmName("step18333")
fun step18333(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
class Counter18334 {
    var current: Int by Delegates.observable(97) { _, _, new -> println(new) }

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
sealed class Outcome18335 {
    data class Success(val payload: String) : Outcome18335()
    data object Failure : Outcome18335()
}

fun describe18335(outcome: Outcome18335): String = when (outcome) {
    is Outcome18335.Success -> outcome.payload
    Outcome18335.Failure -> "failed"
}
