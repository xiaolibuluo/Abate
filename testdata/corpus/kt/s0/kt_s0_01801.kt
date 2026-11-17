package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter14800 {
    var current: Int by Delegates.observable(39) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull14801(): T? = if (size > 1) this[1] else null

fun render14801(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome14802 {
    data class Success(val payload: String) : Outcome14802()
    data object Failure : Outcome14802()
}

fun describe14802(outcome: Outcome14802): String = when (outcome) {
    is Outcome14802.Success -> outcome.payload
    Outcome14802.Failure -> "failed"
}
object Registry14803 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14803 = "registry-60"
    }
}
