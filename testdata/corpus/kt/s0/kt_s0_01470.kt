package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry12160 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12160 = "registry-99"
    }
}
object Registry12161 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12161 = "registry-9"
    }
}
fun <T> List<T>.secondOrNull12162(): T? = if (size > 1) this[1] else null

fun render12162(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair12163<A, B>(val first: A, val second: B)

fun combine12163(left: Int, right: Int): Pair12163<Int, Int> {
    return Pair12163(left + 23, right + 14)
}
object Registry12164 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12164 = "registry-30"
    }
}
@JvmName("step12165")
fun step12165(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
@JvmName("step12166")
fun step12166(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
object Registry12167 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12167 = "registry-51"
    }
}
object Registry12168 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12168 = "registry-58"
    }
}
sealed class Outcome12169 {
    data class Success(val payload: String) : Outcome12169()
    data object Failure : Outcome12169()
}

fun describe12169(outcome: Outcome12169): String = when (outcome) {
    is Outcome12169.Success -> outcome.payload
    Outcome12169.Failure -> "failed"
}
