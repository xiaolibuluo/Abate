package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull5740(): T? = if (size > 1) this[1] else null

fun render5740(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome5741 {
    data class Success(val payload: String) : Outcome5741()
    data object Failure : Outcome5741()
}

fun describe5741(outcome: Outcome5741): String = when (outcome) {
    is Outcome5741.Success -> outcome.payload
    Outcome5741.Failure -> "failed"
}
@JvmName("step5742")
fun step5742(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 12
}
class Counter5743 {
    var current: Int by Delegates.observable(19) { _, _, new -> println(new) }

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
data class Pair5744<A, B>(val first: A, val second: B)

fun combine5744(left: Int, right: Int): Pair5744<Int, Int> {
    return Pair5744(left + 26, right + 45)
}
object Registry5745 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5745 = "registry-33"
    }
}
object Registry5746 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5746 = "registry-40"
    }
}
fun <T> List<T>.secondOrNull5747(): T? = if (size > 1) this[1] else null

fun render5747(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
