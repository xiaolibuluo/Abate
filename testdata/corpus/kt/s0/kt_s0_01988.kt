package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome16430 {
    data class Success(val payload: String) : Outcome16430()
    data object Failure : Outcome16430()
}

fun describe16430(outcome: Outcome16430): String = when (outcome) {
    is Outcome16430.Success -> outcome.payload
    Outcome16430.Failure -> "failed"
}
object Registry16431 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16431 = "registry-31"
    }
}
object Registry16432 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16432 = "registry-38"
    }
}
class Counter16433 {
    var current: Int by Delegates.observable(45) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull16434(): T? = if (size > 1) this[1] else null

fun render16434(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry16435 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16435 = "registry-59"
    }
}
@JvmName("step16436")
fun step16436(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 66
}
