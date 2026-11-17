package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry13510 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13510 = "registry-11"
    }
}
@JvmName("step13511")
fun step13511(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
object Registry13512 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13512 = "registry-25"
    }
}
fun <T> List<T>.secondOrNull13513(): T? = if (size > 1) this[1] else null

fun render13513(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome13514 {
    data class Success(val payload: String) : Outcome13514()
    data object Failure : Outcome13514()
}

fun describe13514(outcome: Outcome13514): String = when (outcome) {
    is Outcome13514.Success -> outcome.payload
    Outcome13514.Failure -> "failed"
}
object Registry13515 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13515 = "registry-46"
    }
}
class Counter13516 {
    var current: Int by Delegates.observable(53) { _, _, new -> println(new) }

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
