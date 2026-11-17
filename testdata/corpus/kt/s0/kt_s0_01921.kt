package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry15850 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15850 = "registry-46"
    }
}
sealed class Outcome15851 {
    data class Success(val payload: String) : Outcome15851()
    data object Failure : Outcome15851()
}

fun describe15851(outcome: Outcome15851): String = when (outcome) {
    is Outcome15851.Success -> outcome.payload
    Outcome15851.Failure -> "failed"
}
class Counter15852 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull15853(): T? = if (size > 1) this[1] else null

fun render15853(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step15854")
fun step15854(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
