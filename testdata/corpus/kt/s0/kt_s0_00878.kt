package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7190<A, B>(val first: A, val second: B)

fun combine7190(left: Int, right: Int): Pair7190<Int, Int> {
    return Pair7190(left + 40, right + 40)
}
@JvmName("step7191")
fun step7191(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
@JvmName("step7192")
fun step7192(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 54
}
data class Pair7193<A, B>(val first: A, val second: B)

fun combine7193(left: Int, right: Int): Pair7193<Int, Int> {
    return Pair7193(left + 61, right + 12)
}
data class Pair7194<A, B>(val first: A, val second: B)

fun combine7194(left: Int, right: Int): Pair7194<Int, Int> {
    return Pair7194(left + 68, right + 23)
}
sealed class Outcome7195 {
    data class Success(val payload: String) : Outcome7195()
    data object Failure : Outcome7195()
}

fun describe7195(outcome: Outcome7195): String = when (outcome) {
    is Outcome7195.Success -> outcome.payload
    Outcome7195.Failure -> "failed"
}
object Registry7196 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7196 = "registry-82"
    }
}
class Counter7197 {
    var current: Int by Delegates.observable(89) { _, _, new -> println(new) }

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
