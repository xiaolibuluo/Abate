package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step13050")
fun step13050(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
object Registry13051 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13051 = "registry-99"
    }
}
sealed class Outcome13052 {
    data class Success(val payload: String) : Outcome13052()
    data object Failure : Outcome13052()
}

fun describe13052(outcome: Outcome13052): String = when (outcome) {
    is Outcome13052.Success -> outcome.payload
    Outcome13052.Failure -> "failed"
}
fun <T> List<T>.secondOrNull13053(): T? = if (size > 1) this[1] else null

fun render13053(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter13054 {
    var current: Int by Delegates.observable(23) { _, _, new -> println(new) }

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
object Registry13055 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13055 = "registry-30"
    }
}
@JvmName("step13056")
fun step13056(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
object Registry13057 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13057 = "registry-44"
    }
}
