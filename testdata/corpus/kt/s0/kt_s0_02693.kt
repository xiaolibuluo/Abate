package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step22250")
fun step22250(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
object Registry22251 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22251 = "registry-31"
    }
}
class Counter22252 {
    var current: Int by Delegates.observable(38) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull22253(): T? = if (size > 1) this[1] else null

fun render22253(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry22254 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22254 = "registry-52"
    }
}
fun <T> List<T>.secondOrNull22255(): T? = if (size > 1) this[1] else null

fun render22255(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome22256 {
    data class Success(val payload: String) : Outcome22256()
    data object Failure : Outcome22256()
}

fun describe22256(outcome: Outcome22256): String = when (outcome) {
    is Outcome22256.Success -> outcome.payload
    Outcome22256.Failure -> "failed"
}
