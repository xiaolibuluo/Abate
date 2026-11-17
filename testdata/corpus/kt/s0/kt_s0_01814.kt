package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14920")
fun step14920(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 98
}
sealed class Outcome14921 {
    data class Success(val payload: String) : Outcome14921()
    data object Failure : Outcome14921()
}

fun describe14921(outcome: Outcome14921): String = when (outcome) {
    is Outcome14921.Success -> outcome.payload
    Outcome14921.Failure -> "failed"
}
@JvmName("step14922")
fun step14922(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
@JvmName("step14923")
fun step14923(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 22
}
fun <T> List<T>.secondOrNull14924(): T? = if (size > 1) this[1] else null

fun render14924(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry14925 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14925 = "registry-36"
    }
}
object Registry14926 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14926 = "registry-43"
    }
}
