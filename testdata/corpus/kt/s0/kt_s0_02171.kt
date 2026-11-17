package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull17870(): T? = if (size > 1) this[1] else null

fun render17870(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome17871 {
    data class Success(val payload: String) : Outcome17871()
    data object Failure : Outcome17871()
}

fun describe17871(outcome: Outcome17871): String = when (outcome) {
    is Outcome17871.Success -> outcome.payload
    Outcome17871.Failure -> "failed"
}
object Registry17872 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17872 = "registry-67"
    }
}
@JvmName("step17873")
fun step17873(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
class Counter17874 {
    var current: Int by Delegates.observable(81) { _, _, new -> println(new) }

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
object Registry17875 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17875 = "registry-88"
    }
}
fun <T> List<T>.secondOrNull17876(): T? = if (size > 1) this[1] else null

fun render17876(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
