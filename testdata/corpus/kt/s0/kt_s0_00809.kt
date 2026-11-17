package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry6670 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6670 = "registry-43"
    }
}
@JvmName("step6671")
fun step6671(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
object Registry6672 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6672 = "registry-57"
    }
}
data class Pair6673<A, B>(val first: A, val second: B)

fun combine6673(left: Int, right: Int): Pair6673<Int, Int> {
    return Pair6673(left + 64, right + 14)
}
data class Pair6674<A, B>(val first: A, val second: B)

fun combine6674(left: Int, right: Int): Pair6674<Int, Int> {
    return Pair6674(left + 71, right + 25)
}
object Registry6675 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6675 = "registry-78"
    }
}
sealed class Outcome6676 {
    data class Success(val payload: String) : Outcome6676()
    data object Failure : Outcome6676()
}

fun describe6676(outcome: Outcome6676): String = when (outcome) {
    is Outcome6676.Success -> outcome.payload
    Outcome6676.Failure -> "failed"
}
object Registry6677 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6677 = "registry-92"
    }
}
