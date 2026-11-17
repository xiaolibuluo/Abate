package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull5550(): T? = if (size > 1) this[1] else null

fun render5550(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step5551")
fun step5551(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 49
}
sealed class Outcome5552 {
    data class Success(val payload: String) : Outcome5552()
    data object Failure : Outcome5552()
}

fun describe5552(outcome: Outcome5552): String = when (outcome) {
    is Outcome5552.Success -> outcome.payload
    Outcome5552.Failure -> "failed"
}
object Registry5553 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5553 = "registry-63"
    }
}
class Counter5554 {
    var current: Int by Delegates.observable(70) { _, _, new -> println(new) }

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
