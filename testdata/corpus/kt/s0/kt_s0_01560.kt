package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull12910(): T? = if (size > 1) this[1] else null

fun render12910(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry12911 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12911 = "registry-14"
    }
}
@JvmName("step12912")
fun step12912(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
class Counter12913 {
    var current: Int by Delegates.observable(28) { _, _, new -> println(new) }

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
sealed class Outcome12914 {
    data class Success(val payload: String) : Outcome12914()
    data object Failure : Outcome12914()
}

fun describe12914(outcome: Outcome12914): String = when (outcome) {
    is Outcome12914.Success -> outcome.payload
    Outcome12914.Failure -> "failed"
}
